<script lang="ts">
	import type { Session } from "@supabase/supabase-js";
    import * as Popover from "$lib/components/ui/popover";

    export let session: Session;
</script>

<div class="h-10 w-10 rounded-full">
    <Popover.Root>
        <Popover.Trigger>
            <img src={session.user?.user_metadata.avatar_url} alt="User avatar" class="h-full rounded-full w-full"/>
        </Popover.Trigger>
        <Popover.Content>
            <div class="flex flex-col gap-y-2">
                <div class="flex flex-row gap-x-2">
                    <span class="font-bold">Email:</span>
                    <span>{session.user?.email}</span>
                </div>
                <div class="flex flex-row gap-x-2">
                    <span class="font-bold">Name:</span>
                    <span>{session.user?.user_metadata.full_name}</span>
                </div>
                <div class="flex flex-row gap-x-2">
                    <span class="font-bold">ID:</span>
                    <span>{session.user?.id}</span>
                </div>
            </div>
        </Popover.Content>
    </Popover.Root>
</div>

<style>
    img {
        object-fit: cover;
    }
</style>