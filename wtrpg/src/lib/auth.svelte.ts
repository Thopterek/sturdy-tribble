const api = 'localhost:8080';

export const user = $state({
	name: '',
	access_token: '',
	email: '',
	password: '',
});

// Testing version it's fine
export async function create_user(c_email: string, c_name: string, c_pass: string) {
	const payload = {
		email: c_email,
		username: c_name,
		password: c_pass
	};
	try {
		const result = await fetch(`${api}/api/login`, {
			method: "POST",
			headers: {
				"Content-Type": "application/json"
			},
			body: JSON.stringify(payload)
		});
		const data = await result.json();
	}
	catch {
	}
}
