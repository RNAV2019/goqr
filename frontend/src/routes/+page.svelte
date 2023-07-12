<script lang="ts">
	import { PUBLIC_DB } from '$env/static/public';
	import Button from '../components/Button.svelte';
	import CreateModal from '../components/CreateModal.svelte';
	import List from '../components/List.svelte';
	import type { createFieldsType } from '../utils/helper';
	import { refetchData } from '../utils/store';
	let showModal: boolean = false;

	function handleClose() {
		showModal = false;
	}

	async function createGoqr(data: createFieldsType) {
		await fetch(`${PUBLIC_DB}goqr`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(data)
		}).then((res) => console.log(res));
		refetchData();
	}
</script>

<main class="flex flex-col gap-6">
	<div class="flex flex-row justify-between">
		<h1 class="text-2xl font-bold flex flex-row items-center gap-2">
			Goqr <div class="w-4 h-1 bg-black dark:bg-white mt-1" />
			The Go QRCode Generator!
		</h1>
		<Button onClickHandler={() => (showModal = true)}>Create</Button>
	</div>
	<List />
	<CreateModal show={showModal} title={'Create Goqr'} on:close={handleClose} {createGoqr} />
</main>
