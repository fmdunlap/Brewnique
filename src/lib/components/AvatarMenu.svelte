<script lang="ts">
	import type { Session } from "@supabase/supabase-js";

    export let session: Session;
</script>

<div class="h-8 w-8 rounded-full">
    <img src={session.user?.user_metadata.avatar_url} alt="User avatar" class="h-full rounded-full w-full"/>
</div>

<style>
    img {
        object-fit: cover;
    }
</style>