import React, { useState, useEffect, useRef } from 'react';
import { useNavigate, useParams } from "react-router-dom";

import { grpc } from '@improbable-eng/grpc-web';
import { getCookie } from 'typescript-cookie'

import { API } from 'cmd/api/grpc/api_pb_service';
import { ListTemplateReq, ListTemplateResp } from 'pkg/entity/dto/template_pb';
import { CreatePCReq } from 'pkg/entity/dto/pc_pb';
import { PC } from 'pkg/entity/pc_pb';
import { Template } from 'pkg/entity/template_pb';

import { ulid, parse } from '../lib/ulid'

import CircularProgress from '@mui/material/CircularProgress';
import Paper from '@mui/material/Paper';
import Avatar from '@mui/material/Avatar';
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
import { CardHeader } from '@mui/material';

export default () => {
	const { roomID } = useParams()

	const navigate = useNavigate()

	// API Room
	const [templates, setTemplates] = useState<{ templates: Template[], loaded: boolean }>({ templates: [], loaded: false })

	const createPC = (req: CreatePCReq) => {
		let md = new grpc.Metadata()
		md.set('token', getCookie('access')!)

		const prom = new Promise<PC>((resolve, reject) => {
			grpc.unary(API.CreatePC, {
				metadata: md,
				request: req,
				host: 'https://api.legacyfactory.com:8082',
				onEnd: res => {
					const { status, statusMessage, headers, message, trailers } = res;
					if (status !== grpc.Code.OK || !message) {
						reject(res)

						return
					}

					resolve(message as PC)
				}
			});
		})

		return prom
	}

	const listTemplates = (req: ListTemplateReq) => {
		let md = new grpc.Metadata()
		md.set('token', getCookie('access')!)

		const prom = new Promise<ListTemplateResp>((resolve, reject) => {
			grpc.unary(API.ListTemplate, {
				metadata: md,
				request: req,
				host: 'https://api.legacyfactory.com:8082',
				onEnd: res => {
					const { status, statusMessage, headers, message, trailers } = res;
					if (status !== grpc.Code.OK || !message) {
						reject(res)

						return
					}

					resolve(message as ListTemplateResp)
				}
			});
		})

		return prom
	}

	const refreshTemplates = () => {
		const req = new ListTemplateReq()
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
	const [rowsPerPage, setRowsPerPage] = React.useState(100);

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

		const req = new CreatePCReq()
		req.setEntitytemplate(selectedTemplate)
		req.setRoomid(parse(roomID!))

		createPC(req).then((result) => {
			console.log('pc created', ulid(result.getId_asU8()))

			navigate('/rooms')
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
					<div onClick={() => { navigate(-1) }}>
						<Button variant="contained" startIcon={<ArrowBack />} color='primary' size='large'>
							Back
						</Button>
					</div>
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
								style={{ minWidth: 20, width: '10%' }}
							>
								ID
							</TableCell>
							<TableCell
								key='name'
								align='left'
								style={{ minWidth: 20, width: '50%' }}
							>
								Name
							</TableCell>
						</TableRow>
					</TableHead>
					<TableBody key='table_template'>
						{!templates.loaded &&
							<TableRow hover role='checkbox' tabIndex={-1} key='loading'>
								<TableCell
									colSpan={2}
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
								const id = template.getEntityid_asU8()
								const sid = ulid(id)
								const name = template.getName()

								console.log('display line template:', sid, name)
								return (
									<TableRow hover role='checkbox' tabIndex={-1} key={sid} selected={selectedTemplate == name} onClick={() => { setselectedTemplate(name) }}>
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
											style={{ minWidth: 20, width: '50%' }}
										>
											<CardHeader
												avatar={<Avatar alt={name} src={'img/assets/' + name + '.png'} />}
												title={name}
											/>
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
