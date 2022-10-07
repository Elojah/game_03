import React, { useState, useEffect } from 'react';

import { grpc } from '@improbable-eng/grpc-web';

import { getCookie } from 'typescript-cookie'

import API from 'cmd/api/grpc/api_pb_service';
import * as dtoentity from 'pkg/entity/dto/entity_pb';

import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import CircularProgress from '@mui/material/CircularProgress';
import New_entity from './new_entity';

export default () => {

	// const [loadingEntities, toggleLoadingEntities] = useState(true)

	// const [emptyEntities, toggleEmptyEntities] = useState(false)

	// const listEntities = (req: dtoentity.ListEntityReq) => {
	// 	let md = new grpc.Metadata()
	// 	md.set('token', getCookie('token')!)

	// 	const prom = new Promise<dtoentity.ListEntityResp>((resolve, reject) => {
	// 		grpc.unary(API.API.ListEntity, {
	// 			metadata: md,
	// 			request: req,
	// 			host: 'http://localhost:8081',
	// 			onEnd: res => {
	// 				const { status, statusMessage, headers, message, trailers } = res;
	// 				if (status !== grpc.Code.OK || !message) {
	// 					reject(res)

	// 					return
	// 				}

	// 				resolve(message as dtoentity.ListEntityResp)
	// 			}
	// 		});
	// 	})

	// 	return prom
	// }

	// let entities = new dtoentity.ListEntityResp()

	// const refreshEntities = () => {
	// 	const req = new dtoentity.ListEntityReq()
	// 	toggleLoadingEntities(true)
	// 	listEntities(req).then((result) => {
	// 		console.log('found ', result.getEntitiesList().length, ' entities')
	// 		toggleEmptyEntities((result.getEntitiesList().length == 0))
	// 		toggleLoadingEntities(false)

	// 		entities = result
	// 	}).catch((err) => {
	// 		console.log(err)
	// 	})
	// }

	// useEffect(() => { refreshEntities() }, [])

	return (
		<Grid container
			spacing={0}
			direction="column"
			alignItems="center"
			justifyContent="center"
			style={{ minHeight: '100vh' }}
		>
			{/* {loadingEntities &&
				<Grid item xs={3}>
					<CircularProgress color='secondary' />
				</Grid>
			}
			{!loadingEntities && emptyEntities &&
				<Grid item xs={1}>
					<New_entity success={refreshEntities} />
				</Grid>
			}
			{!loadingEntities && !emptyEntities &&
				<>{JSON.stringify(entities)}</>
			} */}
		</Grid>
	);
}
