import {Bot, HardDrive, Users, BarChart3 } from "lucide-react"

export const dashboardSections = [
  {
    title: "Stable Diffusion Management",
    description: "Manages the stable diffusion properties",
    icon: Bot,
    href: "/todo",
    id: 0
  },
  {
    title: "AI Prompting",
    description: "Manages the tasks and prompts of deepseek",
    icon: HardDrive,
    href: "/todo",
    id: 1
  },
  ,
  {
    title: "Team",
    description: "Our team up to this point",
    icon: Users,
    href: "/team",
    id: 2
  },
  {
    title: "Analytics",
    description: "Track your progress and performance",
    icon: BarChart3,
    href: "/analytics",
    id: 3
  },
]

export default function TitleCard({ title, description, icon: Icon, href }) {
  return (
    <a
      href={href}
      className="group relative overflow-hidden rounded-xl bg-gray-100 p-6 transition-all duration-300 hover:bg-gray-200 hover:shadow-lg focus:outline-none focus:ring-2 focus:ring-gray-400"
    >
      <div className="absolute inset-0 bg-gradient-to-br from-gray-50 to-gray-200 opacity-50 transition-opacity group-hover:opacity-70" />
      <div className="relative z-10 flex items-center space-x-4">
        <Icon className="h-8 w-8 text-gray-600" />
        <div>
          <h3 className="text-xl font-semibold text-gray-900 mb-1">{title}</h3>
          <p className="text-sm text-gray-600">{description}</p>
        </div>
      </div>
      <div className="absolute inset-0 backdrop-blur-[2px] opacity-0 transition-opacity group-hover:opacity-100" />
    </a>
  )
}