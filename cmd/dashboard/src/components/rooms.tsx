import React, { useState, useEffect } from 'react';

import { grpc } from '@improbable-eng/grpc-web';

import { getCookie } from 'typescript-cookie'

import API from 'cmd/api/grpc/api_pb_service';
import * as dtoroom from 'pkg/room/dto/room_pb';

import { ulid } from '../lib/ulid'

import { Link, LinkProps } from 'react-router-dom';
import CircularProgress from '@mui/material/CircularProgress';
import Paper from '@mui/material/Paper';
import Grid from '@mui/material/Grid';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TablePagination from '@mui/material/TablePagination';
import TableRow from '@mui/material/TableRow';
import AddCircleIcon from '@mui/icons-material/AddCircle';
import IconButton from '@mui/material/IconButton';
import Searchbar from './searchbar';

export default () => {

	// API Room
	const [loadingRooms, toggleLoadingRooms] = useState(true)

	const listRooms = (req: dtoroom.ListRoomReq) => {
		let md = new grpc.Metadata()
		md.set('token', getCookie('token')!)

		const prom = new Promise<dtoroom.ListRoomResp>((resolve, reject) => {
			grpc.unary(API.API.ListRoom, {
				metadata: md,
				request: req,
				host: 'http://localhost:8082',
				onEnd: res => {
					const { status, statusMessage, headers, message, trailers } = res;
					if (status !== grpc.Code.OK || !message) {
						reject(res)

						return
					}

					resolve(message as dtoroom.ListRoomResp)
				}
			});
		})

		return prom
	}

	let rooms = new dtoroom.ListRoomResp()

	const refreshRooms = () => {
		const req = new dtoroom.ListRoomReq()
		req.setSize(100)
		toggleLoadingRooms(true)
		listRooms(req).then((result) => {
			console.log('found ', result.getRoomsList().length, ' rooms')
			toggleLoadingRooms(false)

			rooms = result
		}).catch((err) => {
			console.log(err)
		})
	}

	useEffect(() => { refreshRooms() }, [])

	// Table Room
	const [page, setPage] = React.useState(0);
	const [rowsPerPage, setRowsPerPage] = React.useState(10);

	const handleChangePage = (event: unknown, newPage: number) => {
		setPage(newPage);
	};

	const handleChangeRowsPerPage = (event: React.ChangeEvent<HTMLInputElement>) => {
		setRowsPerPage(+event.target.value);
		setPage(0);
	};


	return (
		<Paper sx={{ width: '100%', overflow: 'hidden' }}>
			<Grid
				container
				spacing={0}
				alignItems="center"
				justifyContent="center"
				style={{ minHeight: '10vh' }}
			>
				<Grid item xs={1}>
					<IconButton color="primary" size='large' {...{ component: Link, to: "/create_room" }}>
						<AddCircleIcon />
					</IconButton>
				</Grid>
				<Grid item xs={11}>
					<Searchbar />
				</Grid>
			</Grid>
			<TableContainer sx={{ maxHeight: 1080 }}>
				<Table stickyHeader aria-label="sticky table">
					<TableHead>
						<TableRow>
							<TableCell
								key='id'
								align='left'
								style={{ minWidth: 20 }}
							>
								ID
							</TableCell>
							<TableCell
								key='name'
								align='left'
								style={{ minWidth: 100 }}
							>
								Name
							</TableCell>
						</TableRow>
					</TableHead>
					<TableBody>
						{loadingRooms &&
							<TableRow hover role="checkbox" tabIndex={-1} key='loading'>
								<TableCell
									key='id_loading'
									align='left'
									style={{ minWidth: 20 }}
								>
									<CircularProgress color='secondary' />
								</TableCell>
								<TableCell
									key='name_loading'
									align='left'
									style={{ minWidth: 20 }}
								>
									<CircularProgress color='secondary' />
								</TableCell>
							</TableRow>

						}
						{!loadingRooms && rooms.getRoomsList()
							.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
							.map((room) => {
								const id = ulid(room.getRoom()?.getId_asU8()!)
								const name = room.getRoom()?.getName()
								return (
									<TableRow hover role="checkbox" tabIndex={-1} key={id}>
										<TableCell
											key={'id_' + id}
											align='left'
											style={{ minWidth: 20 }}
										>
											{id}
										</TableCell>
										<TableCell
											key={'name_' + id}
											align='left'
											style={{ minWidth: 100 }}
										>
											{room.getRoom()?.getName()}
										</TableCell>
									</TableRow>
								);
							})
						}
					</TableBody>
				</Table>
			</TableContainer>
			<TablePagination
				rowsPerPageOptions={[10, 25, 100]}
				component="div"
				count={rooms.getRoomsList().length}
				rowsPerPage={rowsPerPage}
				page={page}
				onPageChange={handleChangePage}
				onRowsPerPageChange={handleChangeRowsPerPage}
			/>
		</Paper >
	);
}
