<script lang="ts">
	import { variables } from '$lib/variables';
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { Badge, Card, CardBody, CardFooter, CardImg, Icon } from 'sveltestrap';

	import '$lib/custom.scss';

	export let favorite = false;
	export let isRead = false;
	export let id = '';
	export let name = '';

	let thumbnailURL = '';
	let viewURL = '';

	onMount(async () => {
		let u = new URL('/browse/thumbnail', variables.basePath);
		u.searchParams.append('name', name);
		thumbnailURL = u.toString();

		u = new URL('/view', $page.url.origin);
		u.searchParams.append('name', name);
		viewURL = u.toString();

		if (favorite || !isRead) {
			classStr = 'border';
		}

		if (favorite) {
			classStr += ' border-pink';
		} else if (!isRead) {
			classStr += ' border-yellow';
		}
	});

	let classStr = '';
</script>

<Card class="{classStr} h-100" {id}>
	<CardImg top src={thumbnailURL} />
	<CardBody>
		<a href={viewURL}>{name}</a>
	</CardBody>
	<CardFooter>
		{#if favorite}
			<Badge class="bg-pink">
				<span><Icon name="star-fill" /> Favorite </span>
			</Badge>
		{:else if !isRead}
			<Badge class="bg-yellow">
				<span><Icon name="lightning-fill" /> New </span>
			</Badge>
		{:else}
			<Badge>
				<span><Icon name="check" /> Read </span>
			</Badge>
		{/if}
	</CardFooter>
</Card>
