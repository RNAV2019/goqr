<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import type { Goqr, updateFieldsType } from '../utils/helper';
	import TemplateModal from './TemplateModal.svelte';
	import Button from './Button.svelte';
	const dispatch = createEventDispatcher();
	export let show: boolean = true;
	const close = () => {
		show = false;
		dispatch('close');
	};
	export let updateGoqr: (data: updateFieldsType) => void;
	export let goqr: Goqr;
	export let title: string;

	let urlInputError: string = '';
	const update = () => {
		if (goqr.url != undefined && goqr.id != undefined) {
			if (goqr.url.length < 1) {
				urlInputError = 'URL was too short (has to be 8 characters long)';
				return;
			}
			let updateFields: updateFieldsType = {
				url: goqr.url,
				id: goqr.id
			};
			updateGoqr(updateFields);
			close();
		}
	};
</script>

{#if show}
	<TemplateModal {title}>
		<div class="flex flex-row gap-1 items-center mb-2">
			<label for="url">URL: </label>
			<input
				type="text"
				name="url"
				id="url"
				bind:value={goqr.url}
				class="flex-grow p-1 px-2 ml-1 border-2 border-primary bg-transparent outline-none focus:outline-none active:outline-none rounded-lg"
			/>
		</div>
		<h4 class="text-red-400 font-bold text-sm">{urlInputError}</h4>
		<div class="flex justify-end gap-3">
			<Button onClickHandler={update}>Update</Button>
			<Button onClickHandler={close}>Close</Button>
		</div>
	</TemplateModal>
{/if}
