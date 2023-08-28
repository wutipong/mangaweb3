<script lang="ts">
    import { onMount } from 'svelte';
	import { Badge, Card, CardBody, CardFooter, CardImg, Icon } from 'sveltestrap';

    export let name = "";
    export let favorite = false;
    export let id = "";
    export let url = "";
    export let thumbnailURL = "";

    let classStr = '';

    onMount(async () => {
		
		if (favorite) {
			classStr = 'border border-pink';
		}
	});

</script>

<Card class="{classStr} h-100" {id}>
	<CardImg top src={thumbnailURL} />
	<CardBody>
		<a href={url}>{name}</a>
	</CardBody>
	<CardFooter>
		{#if favorite}
			<Badge class="bg-pink">
				<span><Icon name="star-fill" /> Favorite </span>
			</Badge>
        {:else}
        <Badge>
            <span><Icon name="tag" />  </span>
        </Badge>
		{/if}
	</CardFooter>
</Card>