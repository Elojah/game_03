import React, { useState, useEffect, useRef } from 'react';
import { useNavigate } from "react-router-dom";

import { grpc } from '@improbable-eng/grpc-web';
import { getCookie } from 'typescript-cookie'

import API from 'cmd/api/grpc/api_pb_service';
import * as dtotemplate from 'pkg/entity/dto/template_pb';
import * as dtopc from 'pkg/entity/dto/pc_pb';
import * as room from 'pkg/room/room_pb';
import * as pc from 'pkg/entity/pc_pb';
import * as template from 'pkg/entity/template_pb';

import { ulid, parse } from '../lib/ulid'

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
import TextField from '@mui/material/TextField';
import IconButton from '@mui/material/IconButton';
import AddIcon from '@mui/icons-material/Add';
import Searchbar from './searchbar';
import { ArrowBack } from '@mui/icons-material';

interface propCreatePC {
	Room: room.R
}

export default (props: propCreatePC) => {

	const navigate = useNavigate()

	// API Room
	const [templates, setTemplates] = useState<{ templates: template.Template[], loaded: boolean }>({ templates: [], loaded: false })

	const createPC = (req: dtopc.CreatePCReq) => {
		let md = new grpc.Metadata()
		md.set('token', getCookie('token')!)

		const prom = new Promise<pc.PC>((resolve, reject) => {
			grpc.unary(API.API.CreatePC, {
				metadata: md,
				request: req,
				host: 'http://localhost:8082',
				onEnd: res => {
					const { status, statusMessage, headers, message, trailers } = res;
					if (status !== grpc.Code.OK || !message) {
						reject(res)

						return
					}

					resolve(message as pc.PC)
				}
			});
		})

		return prom
	}

	const listTemplates = (req: dtotemplate.ListTemplateReq) => {
		let md = new grpc.Metadata()
		md.set('token', getCookie('token')!)

		const prom = new Promise<dtotemplate.ListTemplateResp>((resolve, reject) => {
			grpc.unary(API.API.ListTemplate, {
				metadata: md,
				request: req,
				host: 'http://localhost:8082',
				onEnd: res => {
					const { status, statusMessage, headers, message, trailers } = res;
					if (status !== grpc.Code.OK || !message) {
						reject(res)

						return
					}

					resolve(message as dtotemplate.ListTemplateResp)
				}
			});
		})

		return prom
	}

	const refreshTemplates = () => {
		const req = new dtotemplate.ListTemplateReq()
		req.setAll(true)
		req.setSize(100)
		setTemplates({ templates: [], loaded: false })

		listTemplates(req).then((result) => {
			console.log('found ', result.getTemplatesList().length, ' templates')

			setTemplates({ templates: result.getTemplatesList(), loaded: true })
		}).catch((err) => {
			console.log(err)
		})
	}

	// Table Template
	const [page, setPage] = React.useState(0);
	const [rowsPerPage, setRowsPerPage] = React.useState(10);

	const handleChangePage = (event: unknown, newPage: number) => {
		setPage(newPage);
	};

	const handleChangeRowsPerPage = (event: React.ChangeEvent<HTMLInputElement>) => {
		setRowsPerPage(+event.target.value);
		setPage(0);
	};

	const [selectedTemplate, setselectedTemplate] = React.useState('');

	const pcName = useRef<HTMLInputElement | null>(null);

	const checkCreatePC = () => {
		if (selectedTemplate.length == 0) {
			// TODO: display input error
			return
		}

		const name = pcName.current?.value
		if (!name) {
			// TODO: display input error
			return
		}

		const req = new dtopc.CreatePCReq()
		req.setEntitytemplate(selectedTemplate)
		req.setRoomid(props.Room.getId_asU8())

		createPC(req).then((result) => {
			console.log('pc created', ulid(result.getId_asU8()))

			navigate('/pcs')
		}).catch((err) => {
			console.log(err)
		})
	}

	useEffect(() => { refreshTemplates() }, [])

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
					<IconButton color='primary' size='large' {...{ component: Link, to: '/rooms' }}>
						<ArrowBack />
					</IconButton>
				</Grid>
				<Grid item xs={4}>
					<TextField id='create-pc-name' inputRef={pcName} label='Name' variant='standard'
						placeholder='e.g: Thylius, etc.'
						InputProps={{
							endAdornment:
								<IconButton color='secondary' size='large' onClick={() => { checkCreatePC() }}>
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
					<TableBody key='table_template'>
						{!templates.loaded &&
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
						{templates.loaded && templates.templates
							.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
							.map((template) => {
								const id = template.getId_asU8()
								const sid = ulid(id)
								const name = template.getName()
								const height = template.getHeight()
								const width = template.getWidth()

								console.log('display line template:', sid)
								return (
									<TableRow hover role='checkbox' tabIndex={-1} key={sid} selected={selectedTemplate == id} onClick={() => { setselectedTemplate(id) }}>
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
				count={templates.templates.length}
				rowsPerPage={rowsPerPage}
				page={page}
				onPageChange={handleChangePage}
				onRowsPerPageChange={handleChangeRowsPerPage}
			/>
		</Paper >
	);
}
