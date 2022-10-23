import React, { useState, useEffect } from 'react';

import { grpc } from '@improbable-eng/grpc-web';

import { getCookie } from 'typescript-cookie'

import { API } from 'cmd/api/grpc/api_pb_service';
import * as session from 'pkg/user/session_pb';
import * as dtoroom from 'pkg/room/dto/room_pb';
import * as dtopc from 'pkg/entity/dto/pc_pb';
import * as dtosession from 'pkg/user/dto/session_pb';

import { ulid } from '../lib/ulid'

import { Link, LinkProps } from 'react-router-dom';
import CircularProgress from '@mui/material/CircularProgress';
import Collapse from '@mui/material/Collapse';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Grid from '@mui/material/Grid';
import KeyboardArrowUpIcon from '@mui/icons-material/KeyboardArrowUp';
import GamesIcon from '@mui/icons-material/Games';
import KeyboardArrowDownIcon from '@mui/icons-material/KeyboardArrowDown';
import Table from '@mui/material/Table';
import TableHead from '@mui/material/TableHead';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableRow from '@mui/material/TableRow';
import AddCircleIcon from '@mui/icons-material/AddCircle';
import IconButton from '@mui/material/IconButton';
import Button from '@mui/material/Button';

interface propRoomsRow {
	Room: dtoroom.Room
}

export default (props: propRoomsRow) => {
	const room = props.Room

	const id = room.getRoom()?.getId_asU8()!
	const sid = ulid(id)
	const name = room.getRoom()?.getName()


	// API PC
	const [pcs, setPCs] = useState<{ pcs: dtopc.PC[], loaded: boolean }>({ pcs: [], loaded: false })

	const listPCs = (req: dtopc.ListPCReq) => {
		let md = new grpc.Metadata()
		md.set('token', getCookie('token')!)

		const prom = new Promise<dtopc.ListPCResp>((resolve, reject) => {
			grpc.unary(API.ListPC, {
				metadata: md,
				request: req,
				host: 'https://api.legacyfactory.com:8082',
				onEnd: res => {
					const { status, statusMessage, headers, message, trailers } = res;
					if (status !== grpc.Code.OK || !message) {
						reject(res)

						return
					}

					resolve(message as dtopc.ListPCResp)
				}
			});
		})

		return prom
	}

	const refreshPCs = () => {
		const req = new dtopc.ListPCReq()
		req.setRoomid(id)
		req.setSize(100)
		setPCs({ pcs: [], loaded: false })
		listPCs(req).then((result) => {
			console.log('found ', result.getPcsList().length, ' pcs')

			setPCs({ pcs: result.getPcsList(), loaded: true })
		}).catch((err) => {
			console.log(err)
		})
	}

	const [open, setOpen] = useState(false);

	useEffect(() => { refreshPCs() }, [open]);

	console.log('display line room:', sid)
	return (<>
		<TableRow
			hover role='checkbox' tabIndex={-1}
			key={'data_' + sid}
			sx={{ '& > *': { borderBottom: 'unset' } }}
		>
			<TableCell
				style={{ minWidth: 20, width: '5%' }}
			>
				<IconButton
					aria-label='expand row'
					size='small'
					onClick={() => setOpen(!open)}
				>
					{open ? <KeyboardArrowUpIcon /> : <KeyboardArrowDownIcon />}
				</IconButton>
			</TableCell>
			<TableCell
				key={'id_' + sid}
				align='left'
				style={{ minWidth: 20, width: '10%' }}
			>
				{sid}
			</TableCell>
			<TableCell
				key={'name_' + sid}
				align='left'
				style={{ minWidth: 100, width: '50%' }}
			>
				{name}
			</TableCell>
		</TableRow>
		<TableRow
			key={'collapse_' + sid}
		>
			<TableCell style={{ paddingBottom: 0, paddingTop: 0 }} colSpan={6}>
				<Collapse in={open} timeout='auto' unmountOnExit>
					<Box sx={{ margin: 1 }}>
						<Grid
							container
							spacing={0}
							alignItems="center"
							justifyContent="center"
							style={{ minHeight: '10vh' }}
						>
							<Grid item xs={2}>
								<Button variant="contained" startIcon={<AddCircleIcon />} color='primary' size='large'   {...{ component: Link, to: '/rooms/' + sid + '/create_pc' }}>
									Create
								</Button>
							</Grid>
							<Grid item xs={9}>
								<Typography variant='h6' gutterBottom component='div'>
									Characters
								</Typography>
							</Grid>
						</Grid>
						<Table size='small' aria-label='purchases'>
							<TableHead>
								<TableRow>
									<TableCell
										style={{ minWidth: 20, width: '20%' }}
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
							<TableBody>
								{!pcs.loaded &&
									<TableRow hover role='checkbox' tabIndex={-1}>
										<TableCell
											colSpan={3}
											align='center'
											style={{ minWidth: 20 }}
										>
											<CircularProgress color='secondary' />
										</TableCell>
									</TableRow>

								}
								{pcs.loaded && pcs.pcs
									.map((pc) => {
										const id = pc.getPc()?.getId_asU8()!
										const sid = ulid(id)
										const name = pc.getEntity()?.getName()

										return (
											<TableRow
												hover role='checkbox' tabIndex={-1}
												key={'data_' + sid}
											>
												<TableCell
													style={{ minWidth: 20, width: '20%' }}
												>
													<Button variant="contained" startIcon={<GamesIcon />} color='primary' size='large' href={'https://client.legacyfactory.com:8080?pc_id=' + sid} target='_blank' rel='noreferrer'>
														Play
													</Button>
												</TableCell>
												<TableCell
													key={'id_' + sid}
													align='left'
													style={{ minWidth: 20, width: '10%' }}
												>
													{sid}
												</TableCell>
												<TableCell
													key={'name_' + sid}
													align='left'
													style={{ minWidth: 100, width: '50%' }}
												>
													{name}
												</TableCell>
											</TableRow>
										)
									})
								}
							</TableBody>
						</Table>
					</Box>
				</Collapse>
			</TableCell>
		</TableRow>
	</>);
}
