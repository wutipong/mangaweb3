<script lang="ts">
	import { Button, Icon } from 'sveltestrap';
	import './custom.scss';

	export let isFavorite = false;
</script>

<button
	class="btn"
	class:btn-pink={isFavorite}
	class:btn-outline-pink={!isFavorite}
    on:click
>
	{#if isFavorite}
		<span><Icon name="star" class="mb-3" /> <slot /></span>
	{:else}
		<span><Icon name="star-open" class="mb-3" /> <slot /></span>
	{/if}
</button>
