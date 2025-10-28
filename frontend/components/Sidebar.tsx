import { JSX } from 'react';
import { FaHome } from "react-icons/fa";
import { IoCalendarSharp } from "react-icons/io5";
import './Sidebar.css';

export default function Sidebar() {
    return (
        <div className="fixed top-0 left-0 h-screen w-16 m-0
                        flex flex-col
                        bg-gray-900 text-white shadow-lg">
            <SidebarIcon icon={<FaHome size={28} />} text='Home' />
            <Divider/>
            <SidebarIcon icon={<IoCalendarSharp size={28} />} text='Schedules' />
        </div>
    );
}

const SidebarIcon = ({ icon, text }: { icon: JSX.Element, text: string }) => {
    return (
        <div className='sidebar-icon group'>
            {icon}
            <span className='sidebar-tooltip group-hover:scale-100'>
                {text}
            </span>
        </div>
    );
}

const Divider = () => <hr className="sidebar-divider" />;