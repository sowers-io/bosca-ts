import {Button} from "@/components/ui/button"
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"
import {Input} from "@/components/ui/input"
import {Label} from "@/components/ui/label"
import {onCreateCollection} from "@/api/collections";
import {useState} from "react";

interface AddCollectionProps {
  collection: string | null | undefined
}

export function AddCollection({collection}: AddCollectionProps) {
  const [name, setName] = useState('');

  const onInputChange = (ev: any) => {
    setName(ev.target.value);
  };

  const [open, setOpen] = useState(false);
  const [loading, setLoading] = useState(false);

  function onSubmit() {
    setLoading(true);
  }

  return (
    <Dialog open={open}>
      <a href="#" onClick={() => setOpen(true)} className="focus:outline-none">
        <span className="absolute inset-0" aria-hidden="true"/>
        <span>Add a Collection</span>
        <span aria-hidden="true"> &rarr;</span>
      </a>
      <DialogContent className="sm:max-w-[425px]">
        <form onSubmit={async (ev) => {
          ev.preventDefault()
          ev.stopPropagation()
          setLoading(true)
          try {
            await onCreateCollection(collection ?? undefined, name)
            setOpen(false)
            setName('')
            window.dispatchEvent(new Event('files-changed'))
          } finally {
            setLoading(false)
          }
        }}>
          <DialogHeader>
            <DialogTitle>Add a Collection</DialogTitle>
            <DialogDescription>
              Add a new collection
            </DialogDescription>
          </DialogHeader>
          <div className="grid gap-4 py-4">
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="name" className="text-right">
                Name
              </Label>
              <Input id="name" value={name} onChange={onInputChange} className="col-span-3" autoComplete={"off"}/>
            </div>
          </div>
          <DialogFooter>
            <Button type="submit" disabled={loading}>Add Collection</Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  )
}
