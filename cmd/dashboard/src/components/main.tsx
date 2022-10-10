import React from 'react';

import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Container from '@mui/material/Container';

import Copyright from './copyright';
import Dashboard from './dashboard';
import PCs from './pcs';
import Rooms from './rooms';

import {
	Route,
	Routes,
} from 'react-router-dom';

export default () => {
	return (
		<Box
			component="main"
			sx={{
				flexGrow: 1,
				height: '100vh',
				overflow: 'auto',
			}}
		>
			<Toolbar />
			<Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
				<Routes>
					<Route index element={<Dashboard />} />
					<Route path="rooms" element={<Rooms />} />
					<Route path="pcs" element={<PCs />} />
				</Routes>
			</Container>
			<Copyright sx={{ pt: 4 }} />
		</Box>
	);
}
