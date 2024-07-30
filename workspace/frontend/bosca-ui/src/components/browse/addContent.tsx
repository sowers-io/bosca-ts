import {
  PhotoIcon,
} from '@heroicons/react/24/outline'
import {AddCollection} from "@/components/browse/addCollection";

const items = [
  {
    component: AddCollection,
    description: 'Great for organizing and grouping content.',
    icon: PhotoIcon,
    background: 'bg-green-500',
  },
]

function classNames(...classes: any[]) {
  return classes.filter(Boolean).join(' ')
}

interface AddContentProps {
  collection?: string | null | undefined
}

export default function AddContent({collection}: AddContentProps) {
  return (
    <div>
      <h2 className="text-base font-semibold leading-6 text-gray-900">Content</h2>
      <p className="mt-1 text-sm text-gray-500">
        Get started by selecting a template or start by uploading a file.
      </p>
      <ul role="list" className="mt-6 grid grid-cols-1 gap-6 border-b border-t border-gray-200 py-6 sm:grid-cols-1">
        {items.map((item, itemIdx) => (
          <li key={itemIdx} className="flow-root">
            <div className="relative -m-2 flex items-center space-x-4 rounded-xl p-2 focus-within:ring-2 focus-within:ring-indigo-500 hover:bg-gray-50">
              <div
                className={classNames(
                  item.background,
                  'flex h-16 w-16 flex-shrink-0 items-center justify-center rounded-lg'
                )}
              >
                <item.icon className="h-6 w-6 text-white" aria-hidden="true" />
              </div>
              <div>
                <h3 className="text-sm font-medium text-gray-900">
                  <item.component collection={collection}/>
                </h3>
                <p className="mt-1 text-sm text-gray-500">{item.description}</p>
              </div>
            </div>
          </li>
        ))}
      </ul>
    </div>
  )
}
