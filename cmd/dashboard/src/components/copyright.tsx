import React from 'react';

import Typography from '@mui/material/Typography';
import Link from '@mui/material/Link';

export default (props: any) => {
	return (
		<Typography variant="body2" color="text.secondary" align="center" {...props}>
			{'Copyright © '}
			<Link color="inherit" href="https://spc.com/">
				SPC
			</Link>{' '}
			{new Date().getFullYear()}
			{'.'}
		</Typography>
	);
}
