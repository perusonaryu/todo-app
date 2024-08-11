// since there's no dynamic data here, we can prerender
// it so that it gets served as a static asset in production
export const prerender = false;

import type { PageLoad } from './$types';
import type { Task } from '../types/task';

export const load: PageLoad = async ({ fetch }) => {
	const res = await fetch(`/api/tasks`);
	const { tasks }: { tasks: Task[] } = await res.json();
	return { tasks };
};
