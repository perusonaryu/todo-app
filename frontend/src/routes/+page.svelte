<script lang="ts">
	import type { PageData } from './$types';
	import { enhance } from '$app/forms';

	export let data: PageData;
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section>
	<form method="POST" action="?/register" use:enhance>
		<label>
			タイトル
			<input name="title" type="text" />
		</label>
		<label>
			詳細
			<textarea name="detail"></textarea>
		</label>
		<button>登録</button>
	</form>
	<div class="bg-blue-500 text-white p-4 rounded-lg">Hello, Tailwind CSS in SvelteKit!</div>
	<ul>
		{#each data.tasks as task}
			<li>{task.title}</li>
		{/each}
	</ul>
</section>

<style>
</style>
