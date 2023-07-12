<script lang="ts">
	import { onMount } from 'svelte';
	import { goqr } from '../utils/store';
	import Card from './Card.svelte';
	import type { Goqr } from '../utils/helper';
	import { PUBLIC_DB } from '$env/static/public';

	onMount(async () => {
		const res = await fetch(`${PUBLIC_DB}goqr`);
		const data = (await res.json()) as Goqr[];
		goqr.set(data);
	});

	$: values = $goqr.sort((a, b) => {
		return a.id! - b.id!;
	});
</script>

<div class="grid grid-cols-1 gap-4 xl:grid-cols-3 xl:gap-8">
	{#each values as goqr (goqr.id)}
		<Card {goqr} />
	{/each}
</div>
