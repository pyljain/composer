'use client'

import ReactQuill from 'react-quill';
import 'react-quill/dist/quill.snow.css';
import '@/editor.css'
import EditorToolbar, {modules, formats} from './EditorToolbar';

interface DocumentEditorProps {
  document: string;
  onDocumentChange: (newDocument: string) => void;
  onSelectionChange: (selectedText: string) => void;
}

export default function DocumentEditor({ document, onDocumentChange, onSelectionChange }: DocumentEditorProps) {
  
  const handleEditorChange = (value: string | undefined) => {
    if (value !== undefined) {
      onDocumentChange(value)
    }
  }

  const handleSelectionChange = (range, source, editor) => {
    if (!range || !editor) {
      return;
    }
    const selectedText = editor.getText().substring(range.index - 1, range.index + range.length)
    onSelectionChange(selectedText)
  }

  return (
    <div className="flex-grow p-8 overflow-auto">
      <div className="mx-auto">
        <div className="relative">
          <EditorToolbar />
          <ReactQuill
              theme="snow"
              value={document}
              onChange={handleEditorChange}
              onChangeSelection={handleSelectionChange}
              style={{ height: "calc(100vh - 200px)" }}
              modules={modules}
              formats={formats}
            />
        </div>
      </div>
    </div>
  )
}

