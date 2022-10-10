import * as React from 'react';

import { Link, LinkProps } from 'react-router-dom';

import { styled } from '@mui/material/styles';

import ListItem from '@mui/material/ListItem';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import ListItemText from '@mui/material/ListItemText';
import DashboardIcon from '@mui/icons-material/Dashboard';
import DomainIcon from '@mui/icons-material/Domain';
import StorefrontIcon from '@mui/icons-material/Storefront';

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
					<DomainIcon />
				</ListItemIcon>
				<ListItemText primary="Rooms" />
			</ListItemButton>
			<ListItemButton {...{ component: Link, to: "/pcs" }}>
				<ListItemIcon>
					<StorefrontIcon />
				</ListItemIcon>
				<ListItemText primary="Characters" />
			</ListItemButton>
		</>
	);
}
