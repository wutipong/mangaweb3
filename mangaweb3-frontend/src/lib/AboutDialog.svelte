<script lang="ts">
    import {
        Modal,
        ModalHeader,
        ModalBody,
        ModalFooter,
        Button,
    } from "sveltestrap/src";
	import { variables } from "./variables";

    export let version = "";

    let open = false;
    const toggle = () => (open = !open);

    export function show() {
        open = true;
    }
</script>

<Modal isOpen={open} {toggle}>
    <ModalHeader {toggle}>About</ModalHeader>
    <ModalBody>
        <h5>MangaWeb</h5>
        <h6>Version {version}</h6>
        <p>&copy; 2021-2023 Wutipong Wongsakuldej. All Right Reserved</p>
        <p>Licensed under MIT License</p>
        <p><a href="https://github.com/wutipong/mangaweb3">Homepage</a></p>
        <h6>Using API server: {variables.basePath}</h6>
    </ModalBody>
    <ModalFooter>
        <Button color="secondary" on:click={toggle}>Close</Button>
    </ModalFooter>
</Modal>
