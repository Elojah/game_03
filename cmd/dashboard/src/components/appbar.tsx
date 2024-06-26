import React from 'react';

import { styled } from '@mui/material/styles';
import MuiAppBar, { AppBarProps as MuiAppBarProps } from '@mui/material/AppBar';

import { width } from './drawer';

interface AppBarProps extends MuiAppBarProps {
	open?: boolean;
}
export default styled(
	MuiAppBar, {
	shouldForwardProp: (prop) => prop !== 'open',
})<AppBarProps>(
	({ theme, open }) => ({
		zIndex: theme.zIndex.drawer + 1,
		transition: theme.transitions.create(['width', 'margin'], {
			easing: theme.transitions.easing.sharp,
			duration: theme.transitions.duration.leavingScreen,
		}),
		...(open && {
			marginLeft: width,
			width: `calc(100% - ${width}px)`,
			transition: theme.transitions.create(['width', 'margin'], {
				easing: theme.transitions.easing.sharp,
				duration: theme.transitions.duration.enteringScreen,
			}),
		}),
	})
);
