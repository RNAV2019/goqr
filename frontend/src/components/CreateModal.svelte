<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import type { createFieldsType } from '../utils/helper';
	import TemplateModal from './TemplateModal.svelte';
	import Button from './Button.svelte';

	const dispatch = createEventDispatcher();
	export let show: boolean = true;
	const close = () => {
		show = false;
		url = '';
		urlInputError = '';
		dispatch('close');
	};
	export let createGoqr: (data: createFieldsType) => void;
	export let title: string;
	let url: string;

	let urlInputError: string = '';
	const create = () => {
		if (url != '') {
			if (url.length <= 1) {
				urlInputError = 'Redirect was too short';
				return;
			}
			let createFields: createFieldsType = {
				url: url
			};
			createGoqr(createFields);
			close();
		} else {
			urlInputError = 'Enter a URL to shorten';
		}
	};
</script>

{#if show}
	<TemplateModal {title}>
		<div class="flex flex-row gap-1 items-center">
			<label for="url">URL: </label>
			<input
				type="text"
				name="url"
				id="url"
				bind:value={url}
				class="flex-grow p-1 px-2 ml-1 border-2 border-primary bg-transparent outline-none focus:outline-none active:outline-none rounded-lg"
			/>
		</div>
		<h4 class="text-red-400 font-bold text-sm">{urlInputError}</h4>
		<div class="flex justify-end gap-3">
			<Button onClickHandler={create}>Create</Button>
			<Button onClickHandler={close}>Close</Button>
		</div>
	</TemplateModal>
{/if}
