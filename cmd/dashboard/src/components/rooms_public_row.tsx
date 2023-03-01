import React from 'react';
import { useNavigate } from "react-router-dom";

import { grpc } from '@improbable-eng/grpc-web';

import { getCookie } from 'typescript-cookie'

import { API } from 'cmd/api/grpc/api_pb_service';
import { Room } from 'pkg/room/dto/room_pb';
import { User } from 'pkg/room/user_pb';
import { CreateRoomUserReq } from 'pkg/room/dto/user_pb';

import { ulid } from '../lib/ulid'

import TableCell from '@mui/material/TableCell';
import TableRow from '@mui/material/TableRow';
import AddCircleIcon from '@mui/icons-material/AddCircle';
import Button from '@mui/material/Button';

interface propRoomsRow {
	Room: Room
}

export default (props: propRoomsRow) => {
	const room = props.Room
	const navigate = useNavigate()

	const id = room.getRoom()?.getId_asU8()!
	const sid = ulid(id)
	const name = room.getRoom()?.getName()

	const createUser = (req: CreateRoomUserReq) => {
		let md = new grpc.Metadata()
		md.set('token', getCookie('token')!)

		const prom = new Promise<User>((resolve, reject) => {
			grpc.unary(API.CreateRoomUser, {
				metadata: md,
				request: req,
				host: 'https://api.legacyfactory.com:8082',
				onEnd: res => {
					const { status, statusMessage, headers, message, trailers } = res;
					if (status !== grpc.Code.OK || !message) {
						reject(res)

						return
					}

					resolve(message as User)
				}
			});
		})

		return prom
	}

	const checkCreateUser = () => {
		const req = new CreateRoomUserReq()
		req.setRoomid(id)

		createUser(req).then((result) => {
			console.log('room joined')

			navigate('/rooms')
		}).catch((err) => {
			console.log(err)
		})
	}


	console.log('display line room public:', sid)
	return (<>
		<TableRow
			hover role='checkbox' tabIndex={-1}
			key={'data_' + sid}
			sx={{ '& > *': { borderBottom: 'unset' } }}
		>
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
			<TableCell
				key={'join_' + sid}
				align='left'
				style={{ minWidth: 100, width: '20%' }}
			>
				<Button variant="contained" startIcon={<AddCircleIcon />} color='primary' size='large' onClick={() => { checkCreateUser() }}>
					Join
				</Button>
			</TableCell>
		</TableRow>
	</>);
}
