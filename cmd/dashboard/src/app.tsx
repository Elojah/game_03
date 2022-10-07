import React from 'react'
import { createRoot } from 'react-dom/client';

import { createTheme, ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import Box from '@mui/material/Box';

import './app.scss';

import Layout from './components/layout';

import {
	MemoryRouter,
} from 'react-router-dom';

function render() {

	const root = createRoot(document.getElementById('root')!)

	root.render(
		<>
			<MemoryRouter>
				<ThemeProvider theme={createTheme({
					palette: {
						mode: 'dark',
					},
				})}>
					<Box sx={{ display: 'flex' }}>
						<CssBaseline />
						<Layout />
					</Box>
				</ThemeProvider>
			</MemoryRouter>
		</ >
	);
}

function main() {
	render()
};

main();
