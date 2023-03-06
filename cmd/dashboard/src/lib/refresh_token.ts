let id: number

function refresh() {
	fetch('https://dashboard.legacyfactory.com:8081/refresh_token', {
		method: 'GET',
	})
		.then((resp) => {
			if (!resp.ok) {
				console.log('failed to refresh token:', resp.status)

				return
			}

			console.log('success refresh token:', resp)
		})
		.catch((err) => {
			console.log('failed to refresh token:', err)
		})
}

function start() {
	clearInterval(id)
	id = window.setInterval(refresh, 1000 * 60 * 15)
}

export { start }
