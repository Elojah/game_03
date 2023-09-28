import React from 'react';
import { useNavigate } from "react-router-dom";

import * as grpc from 'grpc-web';

import { getCookie } from 'typescript-cookie'

import { APIClient } from 'cmd/api/grpc/ApiServiceClientPb';
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

	const client = new APIClient('https://api.legacyfactory.com:8082', null)
	const metadata: grpc.Metadata = { 'token': getCookie('access')! }

	const checkCreateUser = () => {
		const req = new CreateRoomUserReq()
		req.setRoomid(id)

		client.createRoomUser(req, metadata).then((result) => {
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
