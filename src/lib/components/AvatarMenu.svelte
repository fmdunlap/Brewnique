<script lang="ts">
	import type { Session } from "@supabase/supabase-js";
    import * as Dropdown from "$lib/components/ui/dropdown-menu";

    export let session: Session;
</script>

<div class="h-10 w-10 rounded-full">
    <Dropdown.Root>
        <Dropdown.Trigger>
            <img src={session.user?.user_metadata.avatar_url} alt="User avatar" class="h-full rounded-full w-full"/>
        </Dropdown.Trigger>
        <Dropdown.Content class="w-72">
            <Dropdown.Label>Profile</Dropdown.Label>
            <Dropdown.Group>
                <Dropdown.Item>Some Item</Dropdown.Item>
                <Dropdown.Item>Another Item</Dropdown.Item>
                <Dropdown.Item>Yet More Thing</Dropdown.Item>
            </Dropdown.Group>
        </Dropdown.Content>
    </Dropdown.Root>
</div>

<style>
    img {
        object-fit: cover;
    }
</style>