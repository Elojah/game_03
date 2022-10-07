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
import AccountBalanceIcon from '@mui/icons-material/AccountBalance';

export default () => {
	return (
		<>
			<ListItemButton {...{ component: Link, to: "/" }}>
				<ListItemIcon>
					<DashboardIcon />
				</ListItemIcon>
				<ListItemText primary="Dashboard" />
			</ListItemButton>
			<ListItemButton {...{ component: Link, to: "/entities" }}>
				<ListItemIcon>
					<DomainIcon />
				</ListItemIcon>
				<ListItemText primary="Entities" />
			</ListItemButton>
			<ListItemButton {...{ component: Link, to: "/assets" }}>
				<ListItemIcon>
					<StorefrontIcon />
				</ListItemIcon>
				<ListItemText primary="Assets" />
			</ListItemButton>
			<ListItemButton {...{ component: Link, to: "/transactions" }}>
				<ListItemIcon>
					<AccountBalanceIcon />
				</ListItemIcon>
				<ListItemText primary="Transactions" />
			</ListItemButton>
		</>
	);
}
