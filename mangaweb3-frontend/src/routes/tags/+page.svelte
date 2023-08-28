<script lang="ts">
	import { page } from '$app/stores';
	import Item from './Item.svelte';
	import Toolbar from './Toolbar.svelte';
	import type { PageData } from './$types';
	import AboutDialog from '$lib/AboutDialog.svelte';
	import MoveToTop from '$lib/MoveToTop.svelte';
	import { getBackendBaseURL } from '$lib/config';
	import {
		Container,
		Spinner,
		Icon,
		Collapse,
		Navbar,
		NavbarToggler,
		NavbarBrand,
		Nav,
		NavItem,
		NavLink,
		Dropdown,
		DropdownToggle,
		DropdownMenu,
		DropdownItem,
		InputGroup,
		InputGroupText,
		Input,
		Button
	} from 'sveltestrap';

	export let data: PageData;

	let favoriteOnly = false;
	let aboutDialog: AboutDialog;

	function toggleFavoriteOnly() {
		favoriteOnly = !favoriteOnly;
	}

	function onAboutClick() {
		aboutDialog.show();
	}

	function createThumbnailUrl(name: string) {
		const output = new URL('/tag/thumbnail', getBackendBaseURL());
		output.searchParams.append('tag', name);

		return output;
	}

	function createBrowseURL(name: string) {
		const output = new URL('/browse', $page.url.origin);
		output.searchParams.append('tag', name);

		return output;
	}
	let navbarToggleOpen = false;
	function handleUpdate(event: CustomEvent<boolean>) {
		navbarToggleOpen = event.detail;
	}

</script>

<!-- Toolbar
	title="Tag List"
	browseURL={new URL('/browse', $page.url.origin).toString()}
	tagListURL={$page.url.toString()}
	onFilterFavorite={toggleFavoriteOnly}
	{favoriteOnly}
	{onAboutClick}
/-->

<Navbar color="dark" dark expand="md" sticky={'top'}>
	<NavbarBrand href="/">{`Tag list`}</NavbarBrand>
	<NavbarToggler on:click={() => (navbarToggleOpen = !navbarToggleOpen)} />
	<Collapse isOpen={navbarToggleOpen} navbar expand="md" on:update={handleUpdate}>
		<Nav navbar>
			<Dropdown nav inNavbar>
				<DropdownToggle nav caret>Browse</DropdownToggle>
				<DropdownMenu end>
					<DropdownItem><Icon name="list-ul" class="me-3" /> All items</DropdownItem>
					<DropdownItem><Icon name="tags-fill" class="me-3" /> Tag list</DropdownItem>
				</DropdownMenu>
			</Dropdown>
			<Dropdown nav inNavbar>
				<DropdownToggle nav caret>Filter</DropdownToggle>
				<DropdownMenu end>
					<DropdownItem><Icon name="star" class="me-3" /> Favorite</DropdownItem>
				</DropdownMenu>
			</Dropdown>
			
			<NavItem>
				<NavLink on:click={() => aboutDialog.show()}>About</NavLink>
			</NavItem>
		</Nav>
	</Collapse>
</Navbar>


<Container fluid style="padding-top:30px;">
	<div class="grid-container">
		{#each data.tags as tag}
			{#if !favoriteOnly || (favoriteOnly && tag.favorite)}
				<Item
					name={tag.name}
					favorite={tag.favorite}
					id={tag.id.toString()}
					url={createBrowseURL(tag.name).toString()}
					thumbnailURL={createThumbnailUrl(tag.name).toString()}
				/>
			{/if}
		{/each}
	</div>
</Container>

<AboutDialog bind:this={aboutDialog} version={'development'} />

<MoveToTop/>