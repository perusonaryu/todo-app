import { fail } from '@sveltejs/kit';
import type { Actions } from './$types';

export const actions: Actions = {
	register: async ({ request, fetch }) => {
		const data = await request.formData();
		const title = data.get('title');
		const detail = data.get('detail');
		console.log('title:', title);
		console.log('detail:', detail);
		try {
			const response = await fetch('/api/create', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ title, detail })
			});

			const result = await response.json();

			if (!response.ok) {
				return fail(response.status, {
					error: true,
					message: result.error || 'エラーが発生しました'
				});
			}

			return {
				success: true,
				message: '正常に作成されました'
			};
		} catch (error) {
			console.error('エラー:', error);
			return fail(500, {
				error: true,
				message: 'サーバーエラーが発生しました'
			});
		}
	}
};
