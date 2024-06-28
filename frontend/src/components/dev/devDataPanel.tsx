import { Card, CardContent, CardHeader } from '@/components/ui/card'

type DevDataPanelProps = {
    title: string,
    children?: React.ReactNode
}


export function DevDataPanel({ title, children }: DevDataPanelProps) {
    return (
        <Card>
            <CardHeader>
                <p className="text-xl">{title}</p>
            </CardHeader>
            <CardContent className="flex flex-col gap-4">
                {children}
            </CardContent>
        </Card>
    )
}