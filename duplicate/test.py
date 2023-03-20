'''
Test file for duplicate.py
'''

import unittest, os, duplicate

class TestDuplicate(unittest.TestCase):
    def test_main(self):
        # prepare input responses
        rename = 'y'
        dir = os.getcwd()
        pattern = '*.jpg'
        times = 2
        # run the main function
        duplicate.rename_files(dir, pattern)
        duplicate.duplicate_files(dir, pattern, times)
        
        # check the results
        if(rename == 'y'):
            # check the files have been renamed
            files = duplicate.get_files(dir, pattern)
            self.assertEqual(len(files), 3 * (times + 1))
            for i, file in enumerate(files):
                # Get the new file number
                new_file_name = '{}.jpg'.format(i)
                # Get the full path of the new file
                new_file_path = os.path.join(dir, new_file_name)
                self.assertEqual(file, new_file_path)
        
        
        