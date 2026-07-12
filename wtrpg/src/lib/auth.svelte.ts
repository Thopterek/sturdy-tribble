const api = 'localhost:8080';

export const user = $state({
	name: '',
	access_token: '',
	email: '',
	password: '',
});

// Testing version it's fine
// I mean I already made a mistake so we gonna see about it I guess
export async function create_user(c_email: string, c_name: string, c_pass: string): Promise<string> {
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
		if (data.token) {
			return data.token;
		}
	}
	catch (error) {
		return "";
	}
	return "";
}

// same here for login
export async function login(l_email: string, l_pass: string) {
	const payload = {
		email: l_email,
		password: l_pass
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
	catch { }
}
