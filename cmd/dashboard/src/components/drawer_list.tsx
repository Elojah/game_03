import * as React from 'react';

import { Link } from 'react-router-dom';

import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import DashboardIcon from '@mui/icons-material/Dashboard';
import StorageIcon from '@mui/icons-material/Storage';
import PersonIcon from '@mui/icons-material/Person';

export default () => {
	return (
		<>
			<ListItemButton {...{ component: Link, to: "/" }}>
				<ListItemIcon>
					<DashboardIcon />
				</ListItemIcon>
				<ListItemText primary="Dashboard" />
			</ListItemButton>
			<ListItemButton {...{ component: Link, to: "/rooms" }}>
				<ListItemIcon>
					<StorageIcon />
				</ListItemIcon>
				<ListItemText primary="Rooms" />
			</ListItemButton>
			<ListItemButton {...{ component: Link, to: "/pcs" }}>
				<ListItemIcon>
					<PersonIcon />
				</ListItemIcon>
				<ListItemText primary="Characters" />
			</ListItemButton>
		</>
	);
}
