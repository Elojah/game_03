import React from 'react';

import { grpc } from '@improbable-eng/grpc-web';

import { getCookie } from 'typescript-cookie'

import API from 'cmd/api/grpc/api_pb_service';
import * as room from 'pkg/room/room_pb';
import * as dtoroom from 'pkg/room/dto/room_pb';

import AddCircleIcon from '@mui/icons-material/AddCircle';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';

type props = {
	success: () => void
}

const NewRoom: React.FC<props> = ({ success }: props) => {
	const createRoom = (req: room.R) => {
		let md = new grpc.Metadata()
		md.set('token', getCookie('token')!)

		const prom = new Promise<room.R>((resolve, reject) => {
			grpc.unary(API.API.CreateRoom, {
				metadata: md,
				request: req,
				host: 'http://localhost:8081',
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

	const onclick = () => {
		const req = new room.R()

		createRoom(req).then((result) => {
			console.log('room created')
			// refresh parent
			success()
		})
	}

	return (
		<IconButton color="primary" size='large' onClick={onclick}>
			<AddCircleIcon />
		</IconButton>
	);
}

export default NewRoom
