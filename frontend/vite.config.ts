import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [sveltekit()],
	server: {
		proxy: {
			'/api': 'http://go-app:8080'
		},
		port: 5173,
		host: true, // この行を追加して、ホスト全体からアクセス可能にする
		strictPort: true
	}
});
