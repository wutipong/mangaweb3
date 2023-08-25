<script lang="ts">
	import { Icon, Button } from 'sveltestrap';
	function moveToTop() {
		location.hash = '#top';
		location.hash = '';
	}
</script>

<nav aria-label="Move to top navigation" class="position-fixed bottom-0 end-0 p-3">
	<Button color="secondary" on:click={moveToTop}>
		<Icon name="chevron-double-up" />
		<span class="d-none d-sm-block">Top</span>
	</Button>
</nav>
