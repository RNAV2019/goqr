import { writable } from 'svelte/store';
import type { Goqr } from './helper';
import { PUBLIC_DB } from '$env/static/public';

export const goqr = writable([] as Goqr[]);

export async function refetchData() {
	const data = await fetch(`${PUBLIC_DB}goqr`);
	const res = (await data.json()) as Goqr[];
	goqr.set(res);
}
