<script lang="ts">
	import type { Session, SupabaseClient } from "@supabase/supabase-js";
    import * as Dropdown from "$lib/components/ui/dropdown-menu";
	import * as Avatar from "$lib/components/ui/avatar";
	import { invalidateAll } from "$app/navigation";

    export let session: Session;
    export let supabase: SupabaseClient;

    const signOut = async () => {
        await supabase.auth.signOut();
        invalidateAll();
    }
</script>

<div class="h-10 w-10 rounded-full">
    <Dropdown.Root>
        <Dropdown.Trigger>
            <Avatar.Root>
                <Avatar.Image src={session.user?.user_metadata.avatar_url} alt="Avatar" />
                <Avatar.Fallback>{session.user?.email?.slice(0,1)}</Avatar.Fallback>
            </Avatar.Root>
        </Dropdown.Trigger>
        <Dropdown.Content class="w-72">
            <Dropdown.Label>Profile</Dropdown.Label>
            <Dropdown.Group>
                <Dropdown.Item>Some Item</Dropdown.Item>
                <Dropdown.Item>Another Item</Dropdown.Item>
                <Dropdown.Item>Yet More Thing</Dropdown.Item>
                <Dropdown.Item on:click={signOut}>Log Out</Dropdown.Item>
            </Dropdown.Group>
        </Dropdown.Content>
    </Dropdown.Root>
</div>