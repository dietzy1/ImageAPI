import { react } from '@babel/types'
import React from 'react'
import {AiOutlineClose, AiOutlineMenu} from 'react-icons/ai'

function Appbar() {
  return (
    <div className= 'text-white flex justify-between items-center h-24 max-w-[1400px] mx-auto px-4'> {/* 1240 */}
  <h1 className='w-full text-3xl font-bold text-[#00df9a]'>
    Image-Gallery
  </h1>
<ul className='flex'>
<li className='p-4'> </li>
<li className='p-4'> Home </li>
<li className='p-4'> About</li>
<li className='p-4'> API-documentation</li>
</ul>


   </div>
  )
}

export default Appbar