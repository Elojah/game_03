import React, { useState, useEffect } from 'react';

import { grpc } from '@improbable-eng/grpc-web';

import { getCookie } from 'typescript-cookie'

import API from 'cmd/api/grpc/api_pb_service';
import * as dtoworld from 'pkg/room/dto/world_pb';
import * as room from 'pkg/room/room_pb';

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
import IconButton from '@mui/material/IconButton';
import Searchbar from './searchbar';
import { ArrowBack } from '@mui/icons-material';

export default () => {

	// API Room
	const [loadingWorld, toggleLoadingWorlds] = useState(true)

	const createRoom = (req: room.R) => {
		let md = new grpc.Metadata()
		md.set('token', getCookie('token')!)

		const prom = new Promise<room.R>((resolve, reject) => {
			grpc.unary(API.API.CreateRoom, {
				metadata: md,
				request: req,
				host: 'http://localhost:8082',
				onEnd: res => {
					const { status, statusMessage, headers, message, trailers } = res;
					if (status !== grpc.Code.OK || !message) {
						reject(res)

						return
					}

					resolve(message as room.R)
				}
			});
		})

		return prom
	}


	const listWorlds = (req: dtoworld.ListWorldReq) => {
		let md = new grpc.Metadata()
		md.set('token', getCookie('token')!)

		const prom = new Promise<dtoworld.ListWorldResp>((resolve, reject) => {
			grpc.unary(API.API.ListWorld, {
				metadata: md,
				request: req,
				host: 'http://localhost:8082',
				onEnd: res => {
					const { status, statusMessage, headers, message, trailers } = res;
					if (status !== grpc.Code.OK || !message) {
						reject(res)

						return
					}

					resolve(message as dtoworld.ListWorldResp)
				}
			});
		})

		return prom
	}

	let worlds = new dtoworld.ListWorldResp()

	const refreshWorlds = () => {
		const req = new dtoworld.ListWorldReq()
		req.setAll(true)
		req.setSize(100)
		toggleLoadingWorlds(true)
		listWorlds(req).then((result) => {
			console.log('found ', result.getWorldsList().length, ' worlds')
			toggleLoadingWorlds(false)

			worlds = result
		}).catch((err) => {
			console.log(err)
		})
	}

	useEffect(() => { refreshWorlds() }, [])

	// Table World
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
					<IconButton color="primary" size='large' {...{ component: Link, to: "/rooms" }}>
						<ArrowBack />
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
								key='height'
								align='left'
								style={{ minWidth: 100 }}
							>
								Height
							</TableCell>
							<TableCell
								key='width'
								align='left'
								style={{ minWidth: 100 }}
							>
								Width
							</TableCell>
						</TableRow>
					</TableHead>
					<TableBody>
						{loadingWorld &&
							<TableRow hover role="checkbox" tabIndex={-1} key='loading'>
								<TableCell
									key='id_loading'
									align='left'
									style={{ minWidth: 20 }}
								>
									<CircularProgress color='secondary' />
								</TableCell>
								<TableCell
									key='height_loading'
									align='left'
									style={{ minWidth: 20 }}
								>
									<CircularProgress color='secondary' />
								</TableCell>
								<TableCell
									key='width_loading'
									align='left'
									style={{ minWidth: 20 }}
								>
									<CircularProgress color='secondary' />
								</TableCell>
							</TableRow>

						}
						{!loadingWorld && worlds.getWorldsList()
							.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
							.map((world) => {
								const id = ulid(world.getId_asU8()!)
								const height = world.getHeight()
								const width = world.getWidth()
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
											key={'height_' + id}
											align='left'
											style={{ minWidth: 100 }}
										>
											{height}
										</TableCell>
										<TableCell
											key={'width_' + id}
											align='left'
											style={{ minWidth: 100 }}
										>
											{width}
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
				count={worlds.getWorldsList().length}
				rowsPerPage={rowsPerPage}
				page={page}
				onPageChange={handleChangePage}
				onRowsPerPageChange={handleChangeRowsPerPage}
			/>
		</Paper >
	);
}
