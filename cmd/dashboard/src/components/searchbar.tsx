import React from 'react';

import TextField from '@mui/material/TextField';
import IconButton from '@mui/material/IconButton';
import SearchIcon from '@mui/icons-material/Search';

export default () => {
	return (
		<form>
			<TextField
				id="search-bar"
				className="text"
				onInput={(e) => { }}
				label="Search"
				variant="outlined"
				placeholder="Search..."
				size="small"
				color='secondary'
			/>
			<IconButton type="submit" aria-label="search">
				<SearchIcon style={{ fill: "blue" }} />
			</IconButton>
		</form>
	);
}

