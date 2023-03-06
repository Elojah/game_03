import React, { useState, useEffect } from 'react';

import { grpc } from '@improbable-eng/grpc-web';

import { getCookie } from 'typescript-cookie'

import { API } from 'cmd/api/grpc/api_pb_service';
import { ListRoomReq, ListRoomResp, Room } from 'pkg/room/dto/room_pb';

import { ulid } from '../lib/ulid'

import { Link } from 'react-router-dom';
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
import Button from '@mui/material/Button';
import Searchbar from './searchbar';
import RoomsPublicRow from './rooms_public_row';

export default () => {

	// API Room
	const [rooms, setRooms] = useState<{ rooms: Room[], loaded: boolean }>({ rooms: [], loaded: false })

	const listPublicRooms = (req: ListRoomReq) => {
		let md = new grpc.Metadata()
		md.set('token', getCookie('access')!)

		const prom = new Promise<ListRoomResp>((resolve, reject) => {
			grpc.unary(API.ListRoomPublic, {
				metadata: md,
				request: req,
				host: 'https://api.legacyfactory.com:8082',
				onEnd: res => {
					const { status, statusMessage, headers, message, trailers } = res;
					if (status !== grpc.Code.OK || !message) {
						reject(res)

						return
					}

					resolve(message as ListRoomResp)
				}
			});
		})

		return prom
	}

	const refreshRooms = () => {
		const req = new ListRoomReq()
		req.setSize(100)
		setRooms({ rooms: [], loaded: false })
		listPublicRooms(req).then((result) => {
			console.log('found ', result.getRoomsList().length, ' rooms')

			setRooms({ rooms: result.getRoomsList(), loaded: true })
		}).catch((err) => {
			console.log(err)
		})
	}

	// Table Room
	const [page, setPage] = React.useState(0);
	const [rowsPerPage, setRowsPerPage] = React.useState(100);

	const handleChangePage = (event: unknown, newPage: number) => {
		setPage(newPage);
	};

	const handleChangeRowsPerPage = (event: React.ChangeEvent<HTMLInputElement>) => {
		setRowsPerPage(+event.target.value);
		setPage(0);
	};

	useEffect(() => { refreshRooms() }, [])

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
					<Button variant="contained" startIcon={<AddCircleIcon />} color='primary' size='large'  {...{ component: Link, to: "/create_room" }}>
						Create
					</Button>
				</Grid>
				<Grid item xs={10}>
					<Searchbar />
				</Grid>
			</Grid>
			<TableContainer sx={{ maxHeight: 1080 }}>
				<Table stickyHeader aria-label="sticky table">
					<TableHead>
						<TableRow>
							<TableCell
								key='collapse'
								style={{ minWidth: 20, width: '5%' }}
							/>
							<TableCell
								key='id'
								align='left'
								style={{ minWidth: 20, width: '10%' }}
							>
								ID
							</TableCell>
							<TableCell
								key='name'
								align='left'
								style={{ minWidth: 100, width: '50%' }}
							>
								Name
							</TableCell>
						</TableRow>
					</TableHead>
					<TableBody key='table_room'>
						{!rooms.loaded &&
							<TableRow hover role="checkbox" tabIndex={-1}>
								<TableCell
									colSpan={3}
									align='center'
									style={{ minWidth: 20 }}
								>
									<CircularProgress color='secondary' />
								</TableCell>
							</TableRow>
						}
						{rooms.loaded && rooms.rooms
							.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
							.map((room) => {
								return (<RoomsPublicRow Room={room} key={ulid(room.getRoom()?.getId_asU8()!)} />)
							})
						}
					</TableBody>
				</Table>
			</TableContainer>
			<TablePagination
				rowsPerPageOptions={[10, 25, 100]}
				component="div"
				count={rooms.rooms.length}
				rowsPerPage={rowsPerPage}
				page={page}
				onPageChange={handleChangePage}
				onRowsPerPageChange={handleChangeRowsPerPage}
			/>
		</Paper >
	);
}
