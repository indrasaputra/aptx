# Development Guide

- Fork or clone the project

    ```
    $ git@github.com:indrasaputra/url-shortener.git
    ```

- Create a meaningful branch

    ```
    $ git checkout -b <your-meaningful-branch>
    ```

    e.g:

    ```
    $ git checkout -b optimize-stream-url
    ```

- Create some changes and their tests (unit test and any test if any).

- If you want to generate files based on protocol buffer definition, run

    ```
    $ make gengrpc
    ```

- If you want to generate mock based on some interfaces, run

    ```
    $ make mockgen
    ```

- Make sure you format/beautify the code by running

    ```
    $ make pretty
    ```

- As a reminder, always run the command above before add and commit changes.
    That command will be run in CI Pipeline to verify the format.

- Add, commit, and push the changes to repository

    ```
    $ git add .
    $ git commit -s -m "your meaningful commit message"
    $ git push origin <your-meaningful-branch>
    ```

- Create a Pull Request (PR). In your PR's description, please explain the goal of the PR and its changes.

- Ask the other contributors to review.

- Once your PR is approved and its pipeline status is green, ask the owner to merge your PR.