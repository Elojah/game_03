import React, { useState, useEffect } from 'react';

import { grpc } from '@improbable-eng/grpc-web';

import { getCookie } from 'typescript-cookie'

import API from 'cmd/api/grpc/api_pb_service';
import * as dtoroom from 'pkg/room/dto/room_pb';

import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import CircularProgress from '@mui/material/CircularProgress';
import New_room from './new_room';

export default () => {

	const [loadingRooms, toggleLoadingRooms] = useState(true)

	const [emptyRooms, toggleEmptyRooms] = useState(false)

	const listRooms = (req: dtoroom.ListRoomReq) => {
		let md = new grpc.Metadata()
		md.set('token', getCookie('token')!)

		const prom = new Promise<dtoroom.ListRoomResp>((resolve, reject) => {
			grpc.unary(API.API.ListRoom, {
				metadata: md,
				request: req,
				host: 'http://localhost:8082',
				onEnd: res => {
					const { status, statusMessage, headers, message, trailers } = res;
					if (status !== grpc.Code.OK || !message) {
						reject(res)

						return
					}

					resolve(message as dtoroom.ListRoomResp)
				}
			});
		})

		return prom
	}

	let rooms = new dtoroom.ListRoomResp()

	const refreshRooms = () => {
		const req = new dtoroom.ListRoomReq()
		toggleLoadingRooms(true)
		listRooms(req).then((result) => {
			console.log('found ', result.getRoomsList().length, ' rooms')
			toggleEmptyRooms((result.getRoomsList().length == 0))
			toggleLoadingRooms(false)

			rooms = result
		}).catch((err) => {
			console.log(err)
		})
	}

	useEffect(() => { refreshRooms() }, [])

	return (
		<Grid container
			spacing={0}
			direction="column"
			alignItems="center"
			justifyContent="center"
			style={{ minHeight: '100vh' }}
		>
			{loadingRooms &&
				<Grid item xs={3}>
					<CircularProgress color='secondary' />
				</Grid>
			}
			{!loadingRooms && emptyRooms &&
				<Grid item xs={1}>
					<New_room success={refreshRooms} />
				</Grid>
			}
			{!loadingRooms && !emptyRooms &&
				<>{JSON.stringify(rooms)}</>
			}
		</Grid>
	);
}
