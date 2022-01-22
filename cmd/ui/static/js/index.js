let backend = 'http://localhost:9001'

function deleteReq(id) {
	fetch(`${backend}/books/${id}`, {
		method: 'DELETE',
	}).then(() => {
		window.location.reload()
	}).catch((e) => {
		console.log(`Error: ${e}`)
	})
}
