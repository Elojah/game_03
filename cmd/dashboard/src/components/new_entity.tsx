import React from 'react';

import { grpc } from '@improbable-eng/grpc-web';

import { getCookie } from 'typescript-cookie'

import API from 'cmd/api/grpc/api_pb_service';
import * as entity from 'pkg/entity/entity_pb';
import * as dtoentity from 'pkg/entity/dto/entity_pb';

import AddCircleIcon from '@mui/icons-material/AddCircle';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';

type props = {
	success: () => void
}

const NewEntity: React.FC<props> = ({ success }: props) => {
	// const createEntity = (req: dtoentity.CreateEntityReq) => {
	// 	let md = new grpc.Metadata()
	// 	md.set('token', getCookie('token')!)

	// 	const prom = new Promise<entity.E>((resolve, reject) => {
	// 		grpc.unary(API.API.CreateEntity, {
	// 			metadata: md,
	// 			request: req,
	// 			host: 'http://localhost:8081',
	// 			onEnd: res => {
	// 				const { status, statusMessage, headers, message, trailers } = res;
	// 				if (status !== grpc.Code.OK || !message) {
	// 					reject(res)

	// 					return
	// 				}

	// 				resolve(message as entity.E)
	// 			}
	// 		});
	// 	})

	// 	return prom
	// }

	// const onclick = () => {
	// 	const req = new dtoentity.CreateEntityReq()

	// 	createEntity(req).then((result) => {
	// 		console.log('entity created')
	// 		// refresh parent
	// 		success()
	// 	})
	// }

	return (
		<>
			<Typography variant="h6">Create new entity</Typography>
			<IconButton color="primary" size='large' onClick={onclick}>
				<AddCircleIcon />
			</IconButton>
		</>
	);
}

export default NewEntity
