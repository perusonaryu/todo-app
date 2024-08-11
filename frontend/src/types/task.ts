export type Task = {
	id: string;
	title: string;
	detail: string;
	status: 'wait' | 'running' | 'finished';
};
