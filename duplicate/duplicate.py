'''
This is a simple script that will duplicate and rename files in a directory.
'''
import os
import shutil
import re
from glob import glob


def get_files(dir: str, pattern: str) -> list:
    '''Get the list of files in a directory that match a pattern.
    
    Parameters
    ----------
    dir: str
        the directory to search
    pattern: str
        the pattern to match
        
    Returns
    -------
    a list of file paths that match the pattern
    
    '''
    _glob = os.path.join(dir, pattern)
    files = glob(_glob)
    return files


def rename_files(dir :str, pattern :str):
    '''Rename the files in a directory.
    
    Rename the files in a directory to a sequential number starting at 0.
    
    Parameters
    ----------
    dir: str
        the directory to rename the files in
    pattern: str
        A unix glob pattern to match the files to rename.
    '''
    # Get the list of files
    files = get_files(dir, pattern)

    for i, file in enumerate(files):
        # Get the new file number
        new_file_name = '{}.jpg'.format(i)
        # Get the full path of the new file
        new_file_path = os.path.join(dir, new_file_name)

        # Rename the file
        os.rename(file, new_file_path)
        # Update the list of files
        files[i] = new_file_path


def duplicate_files(dir :str, pattern :str, times :int):
    '''Duplicate the files in a directory.
    
    Copy the files in a directory to a new file with a sequential number beginning at the number of files in the directory + 1
    For as many times as specified.
    
    Parameters
    ----------
    dir: str
        the directory to duplicate the files in
    pattern: str
        A unix glob pattern to match the files to duplicate.
    times: int
        The number of times to duplicate the files.
    '''    
    files = get_files(dir, pattern)
    num_files = len(files)
    for n in range(1, times + 1):
        for i, file in enumerate(files):
            # Get the new file number
            idx = num_files * n + i
            new_file_name = '{}.jpg'.format(idx)
            # Get the full path of the new file
            new_file_path = os.path.join(dir, new_file_name)

            # Copy the file
            shutil.copy(file, new_file_path)


def main():
    # Get the current working directory
    dir = input('Enter the directory to duplicate: ')
    pattern = input('Enter the file pattern to duplicate: ')
    times = int(input('Enter the number of times to duplicate: '))
    rename = input('Rename the files? (y/n): ')

    # rename the files if necessary
    if rename == 'y':
        rename_files()

    # duplicate the files
    duplicate_files(dir, pattern, times)
