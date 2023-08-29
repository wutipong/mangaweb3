<script lang="ts">
	import { page } from '$app/stores';
	import Item from './Item.svelte';
	import type { PageData } from './$types';
	import AboutDialog from '$lib/AboutDialog.svelte';
	import MoveToTop from '$lib/MoveToTop.svelte';
	import { variables } from '$lib/variables';
	import {
		Container,
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
		DropdownItem
	} from 'sveltestrap';
	import Toast from '$lib/Toast.svelte';
	import Pagination from '$lib/Pagination.svelte';
	import { goto } from '$app/navigation';

	export let data: PageData;

	$: current_page = data.page;
	$: tags = data.tags;
	$: total_page = data.total_page;

	let favoriteOnly = false;
	let aboutDialog: AboutDialog;
	let toast: Toast;

	function createBrowseURL(name: string) {
		const output = new URL('/browse', $page.url.origin);
		output.searchParams.append('tag', name);

		return output;
	}
	let navbarToggleOpen = false;
	function handleUpdate(event: CustomEvent<boolean>) {
		navbarToggleOpen = event.detail;
	}

	async function recreateThumbnails() {
		const url = new URL('/tag/recreate_thumbnails', variables.basePath);
		await fetch(url);
		toast.show(
			'Re-create thumbnail',
			'Thumbnails recreating in progress. Please refresh after a few minutes.'
		);
	}

	function onPageClick(i: number) {
		let url = new URL($page.url);
		url.searchParams.set('page', i.toString());

		goto(url.toString());
	}
</script>

<Navbar color="dark" dark expand="md" sticky={'top'}>
	<NavbarBrand href="/">{`Tag list`}</NavbarBrand>
	<NavbarToggler on:click={() => (navbarToggleOpen = !navbarToggleOpen)} />
	<Collapse isOpen={navbarToggleOpen} navbar expand="md" on:update={handleUpdate}>
		<Nav navbar>
			<Dropdown nav inNavbar>
				<DropdownToggle nav caret>Browse</DropdownToggle>
				<DropdownMenu end>
					<DropdownItem
						on:click={() => {
							goto(createBrowseURL(''));
						}}
					>
						<Icon name="list-ul" class="me-3" />
						All items
					</DropdownItem>
					<DropdownItem on:click={() => goto($page.url)}>
						<Icon name="tags-fill" class="me-3" />
						Tag list
					</DropdownItem>
				</DropdownMenu>
			</Dropdown>
			<Dropdown nav inNavbar>
				<DropdownToggle nav caret>Filter</DropdownToggle>
				<DropdownMenu>
					<DropdownItem on:click={() => (favoriteOnly = !favoriteOnly)} active={favoriteOnly}>
						<Icon name="star" class="me-3" />
						Favorite
					</DropdownItem>
				</DropdownMenu>
			</Dropdown>
			<Dropdown nav inNavbar>
				<DropdownToggle nav caret>Tools</DropdownToggle>
				<DropdownMenu>
					<DropdownItem on:click={() => recreateThumbnails()}>
						<Icon name="file-image" class="me-3" /> Recreate thumbnails
					</DropdownItem>
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
		<div class="row row-cols-1 row-cols-md-3 row-cols-lg-5 g-3">
			{#each tags as tag}
				{#if !favoriteOnly || (favoriteOnly && tag.favorite)}
					<div class="col">
						<Item {tag} />
					</div>
				{/if}
			{/each}
		</div>
	</div>
</Container>

<div aria-label="Page navigation" class="position-fixed bottom-0 start-50 p-3 translate-middle-x">
	<Pagination bind:currentPage={current_page} bind:totalPage={total_page} />
</div>

<AboutDialog bind:this={aboutDialog} version={'development'} />

<MoveToTop />

<Toast bind:this={toast} />
