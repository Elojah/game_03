import React, { useState, useEffect, useRef } from 'react';
import { useNavigate } from "react-router-dom";

import { grpc } from '@improbable-eng/grpc-web';
import { getCookie } from 'typescript-cookie'

import { API } from 'cmd/api/grpc/api_pb_service';
import { ListWorldReq, ListWorldResp } from 'pkg/room/dto/world_pb';
import { R } from 'pkg/room/room_pb';
import { World } from 'pkg/room/world_pb';

import { ulid } from '../lib/ulid'

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
import TextField from '@mui/material/TextField';
import IconButton from '@mui/material/IconButton';
import Button from '@mui/material/Button';
import AddIcon from '@mui/icons-material/Add';
import Searchbar from './searchbar';
import { ArrowBack } from '@mui/icons-material';

export default () => {

	const navigate = useNavigate()

	// API Room
	const [worlds, setWorlds] = useState<{ worlds: World[], loaded: boolean }>({ worlds: [], loaded: false })

	const createRoom = (req: R) => {
		let md = new grpc.Metadata()
		md.set('token', getCookie('access')!)

		const prom = new Promise<R>((resolve, reject) => {
			grpc.unary(API.CreateRoom, {
				metadata: md,
				request: req,
				host: 'https://api.legacyfactory.com:8082',
				onEnd: res => {
					const { status, statusMessage, headers, message, trailers } = res;
					if (status !== grpc.Code.OK || !message) {
						reject(res)

						return
					}

					resolve(message as R)
				}
			});
		})

		return prom
	}


	const listWorlds = (req: ListWorldReq) => {
		let md = new grpc.Metadata()
		md.set('token', getCookie('access')!)

		const prom = new Promise<ListWorldResp>((resolve, reject) => {
			grpc.unary(API.ListWorld, {
				metadata: md,
				request: req,
				host: 'https://api.legacyfactory.com:8082',
				onEnd: res => {
					const { status, statusMessage, headers, message, trailers } = res;
					if (status !== grpc.Code.OK || !message) {
						reject(res)

						return
					}

					resolve(message as ListWorldResp)
				}
			});
		})

		return prom
	}

	const refreshWorlds = () => {
		const req = new ListWorldReq()
		req.setAll(true)
		req.setSize(100)
		setWorlds({ worlds: [], loaded: false })

		listWorlds(req).then((result) => {
			console.log('found ', result.getWorldsList().length, ' worlds')

			setWorlds({ worlds: result.getWorldsList(), loaded: true })
		}).catch((err) => {
			console.log(err)
		})
	}

	// Table World
	const [page, setPage] = React.useState(0);
	const [rowsPerPage, setRowsPerPage] = React.useState(100);

	const handleChangePage = (event: unknown, newPage: number) => {
		setPage(newPage);
	};

	const handleChangeRowsPerPage = (event: React.ChangeEvent<HTMLInputElement>) => {
		setRowsPerPage(+event.target.value);
		setPage(0);
	};

	const [selectedWorld, setselectedWorld] = React.useState(new Uint8Array);

	const roomName = useRef<HTMLInputElement | null>(null);

	const checkCreateRoom = () => {
		if (selectedWorld.length == 0) {
			// TODO: display input error
			return
		}

		const name = roomName.current?.value
		if (!name) {
			// TODO: display input error
			return
		}

		const req = new R()
		req.setName(name)
		req.setWorldid(selectedWorld)

		createRoom(req).then((result) => {
			console.log('room created', ulid(result.getId_asU8()))

			navigate('/rooms')
		}).catch((err) => {
			console.log(err)
		})
	}

	useEffect(() => { refreshWorlds() }, [])

	return (
		<Paper sx={{ width: '100%', overflow: 'hidden' }}>
			<Grid
				container
				spacing={0}
				alignItems='center'
				justifyContent='center'
				style={{ minHeight: '10vh' }}
			>
				<Grid item xs={1}>
					<div onClick={() => { navigate(-1) }}>
						<Button variant="contained" startIcon={<ArrowBack />} color='primary' size='large'>
							Back
						</Button>
					</div>
				</Grid>
				<Grid item xs={4}>
					<TextField id='create-room-name' inputRef={roomName} label='Name' variant='standard'
						placeholder='e.g: Battle for Rohan, etc.'
						InputProps={{
							endAdornment:
								<IconButton color='secondary' size='large' onClick={() => { checkCreateRoom() }}>
									<AddIcon />
								</IconButton>
						}}
					/>
				</Grid>
				<Grid item xs={6}>
					<Searchbar />
				</Grid>
			</Grid>
			<TableContainer sx={{ maxHeight: 1080 }}>
				<Table stickyHeader aria-label='sticky table'>
					<TableHead>
						<TableRow>
							<TableCell
								key='id'
								align='left'
								style={{ minWidth: 20, width: '20%' }}
							>
								ID
							</TableCell>
							<TableCell
								key='name'
								align='left'
								style={{ minWidth: 20, width: '40%' }}
							>
								Name
							</TableCell>
							<TableCell
								key='height'
								align='left'
								style={{ minWidth: 20, width: '10%' }}
							>
								Height
							</TableCell>
							<TableCell
								key='width'
								align='left'
								style={{ minWidth: 20, width: '10%' }}
							>
								Width
							</TableCell>
						</TableRow>
					</TableHead>
					<TableBody key='table_world'>
						{!worlds.loaded &&
							<TableRow hover role='checkbox' tabIndex={-1} key='loading'>
								<TableCell
									colSpan={4}
									key='loading'
									align='center'
									style={{ minWidth: 20 }}
								>
									<CircularProgress color='secondary' />
								</TableCell>
							</TableRow>

						}
						{worlds.loaded && worlds.worlds
							.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
							.map((world) => {
								const id = world.getId_asU8()
								const sid = ulid(id)
								const name = world.getName()
								const height = world.getHeight()
								const width = world.getWidth()

								console.log('display line world:', sid)
								return (
									<TableRow hover role='checkbox' tabIndex={-1} key={sid} selected={selectedWorld == id} onClick={() => { setselectedWorld(id) }}>
										<TableCell
											key={'id_' + sid}
											align='left'
											style={{ minWidth: 20, width: '20%' }}
										>
											{sid}
										</TableCell>
										<TableCell
											key={'name_' + sid}
											align='left'
											style={{ minWidth: 20, width: '40%' }}
										>
											{name}
										</TableCell>
										<TableCell
											key={'height_' + sid}
											align='left'
											style={{ minWidth: 20, width: '10%' }}
										>
											{height}
										</TableCell>
										<TableCell
											key={'width_' + sid}
											align='left'
											style={{ minWidth: 20, width: '10%' }}
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
				component='div'
				count={worlds.worlds.length}
				rowsPerPage={rowsPerPage}
				page={page}
				onPageChange={handleChangePage}
				onRowsPerPageChange={handleChangeRowsPerPage}
			/>
		</Paper >
	);
}
