import React from 'react';

import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Container from '@mui/material/Container';

import Copyright from './copyright';
import Dashboard from './dashboard';
import Assets from './assets';
import Entities from './entities';
import Transactions from './transactions';

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
					<Route path="entities" element={<Entities />} />
					<Route path="assets" element={<Assets />} />
					<Route path="transactions" element={<Transactions />} />
				</Routes>
			</Container>
			<Copyright sx={{ pt: 4 }} />
		</Box>
	);
}
