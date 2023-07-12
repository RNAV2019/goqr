<script lang="ts">
	import { PUBLIC_DB } from '$env/static/public';
	import { Globe, Clipboard } from 'lucide-svelte';
	import type { Goqr, updateFieldsType } from '../utils/helper';
	import { refetchData } from '../utils/store';
	import Button from './Button.svelte';
	import UpdateModal from './UpdateModal.svelte';
	export let goqr: Goqr;
	let showModal: boolean = false;
	$: image = 'data:image/png;base64,' + goqr.qrcode;

	function handleClose() {
		showModal = false;
	}

	async function updateGoqr(data: updateFieldsType) {
		await fetch(`${PUBLIC_DB}Goqr`, {
			method: 'PATCH',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify(data)
		}).then((res) => console.log(res));
		refetchData();
	}

	async function deleteGoqr(id: number | undefined) {
		if (id === undefined) {
			return;
		}
		await fetch(`${PUBLIC_DB}Goqr/${id}`, {
			method: 'DELETE'
		}).then((res) => console.log(res));
		refetchData();
	}

	async function copyImage() {
		const data = await fetch(image);
		const blob = await data.blob();
		await navigator.clipboard.write([
			new ClipboardItem({
				'image/png': blob
			})
		]);
	}
</script>

<div
	class="h-52 rounded-lg bg-gray-100 text-black dark:bg-secondary dark:text-white p-5 flex flex-col justify-between w-full relative"
>
	<div class="flex flex-row gap-4">
		<!-- svelte-ignore a11y-click-events-have-key-events -->
		<!-- svelte-ignore a11y-no-static-element-interactions -->
		<div
			class="relative xl:w-24 xl:h-24 w-44 h-44 flex-shrink-0 group transition-all"
			on:click={() => copyImage()}
		>
			<div
				class="opacity-0 group-hover:opacity-100 absolute top-0 left-0 h-full w-full bg-black/60 flex items-center justify-center cursor-pointer"
			>
				<Clipboard size={22} color="white" />
			</div>
			<img src={image} alt="QRCode" />
		</div>
		<div class="flex flex-row gap-2 items-center self-start mt-1 overflow-hidden">
			<Globe size={16} class="flex-shrink-0" />
			<a
				href={goqr.url}
				target="_blank"
				class="text-base text-ellipsis whitespace-nowrap overflow-hidden hover:underline hover:cursor-pointer flex-grow"
			>
				{goqr.url}
			</a>
		</div>
	</div>
	<div class="absolute right-5 bottom-5 flex flex-row gap-3 self-end">
		<Button onClickHandler={() => (showModal = true)}>Update</Button>
		<Button onClickHandler={() => deleteGoqr(goqr.id)}>Delete</Button>
	</div>
</div>
<UpdateModal show={showModal} title={'Update Goqr'} on:close={handleClose} {updateGoqr} {goqr} />
