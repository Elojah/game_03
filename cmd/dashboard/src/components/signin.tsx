import React from 'react';

import {
	GoogleOAuthProvider,
	GoogleLogin,
	CredentialResponse,
	googleLogout,
} from '@react-oauth/google';

export default () => {
	return (
		<GoogleOAuthProvider clientId={process.env.GOOGLE_CLIENT_ID!}>
			<GoogleLogin
				useOneTap
				auto_select
				theme='filled_black'
				size='large'
				shape='pill'
				width='256'
				onSuccess={GSigninSuccess}
				onError={GSigninError}
			/>
		</GoogleOAuthProvider>
	);
}

function GSigninSuccess(resp: CredentialResponse) {
	fetch('http://localhost:8080/signin', {
		method: 'POST',
		body: resp.credential,
	})
		.then((resp) => {
			if (!resp.ok) {
				console.log('failed to signin:', resp.status)
				googleLogout();

				return
			}

			console.log('success signin:', resp)
		})
		.catch((err) => {
			console.log('failed to signin:', err)
			googleLogout();
		})
}

function GSigninError() {
	console.log('failed to google signin')
}
