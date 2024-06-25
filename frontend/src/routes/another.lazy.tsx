import { createLazyFileRoute } from '@tanstack/react-router'
import {Button} from "@/components/ui/button.tsx";

export const Route = createLazyFileRoute('/another')({
  component: () => <>
      <p>Hello /another!</p>
      <Button onClick={() => alert('clicked')}>Click me</Button>
    </>
})