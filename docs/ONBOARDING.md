
# GitHub Repo Fork

Each challenge member has been invited to a team. Every team has read access to the following repositories:

- [ArubaCloud CLI](https://github.com/Arubacloud/bud-hackaton24-arubacloud-cli)
- [ArubaCloud K8s Autoscaler](https://github.com/Arubacloud/bud-hackaton24-arubacloud-k8s-autoscaler)
- [ArubaCloud SDK](https://github.com/Arubacloud/bud-hackaton24-arubacloud-sdk)
- [ArubaCloud Terraform Provider](https://github.com/Arubacloud/bud-hackaton24-arubacloud-terraform-provider)

To work on a team-maintained repository, one member from each team will need to fork the template repository by following these steps:

1. Click on **Create a new fork**  
   ![Fork Creation Step](./images/Immagine%202024-09-06%20172641.jpg)
   
2. Change the forked repository owner to **ArubaCloud** and update the name to `bud-hackaton24-arubacloud-cli-teamXX`, making sure to use the correct team number.

3. Create the fork  
   ![Forking Process Screenshot](./images/Screenshot%202024-09-06%20172906.jpg)

4. After landing on the forked repository, you will need to grant admin access to your team members:
   - Go to **Settings** > **Collaborators and teams** > **Add teams**
   - Add the `ArubaCloud/team_xx` you are in, with the role set to **admin**  
   ![Team Settings Screenshot](./images/Screenshot%202024-09-06%20173607.jpg)

Your team is now ready to clone, pull, and push to your repository!

---

# GitHub Repo Structure

```
README.md
.gitignore
```

```
src/
    main.**
    utils.**
    config.**
```

```
data/
    kaas.csv
    cloudserver.csv
```

```
docs/
    index.md
    installation.md
```

```
tests/
    test_main.**
    test_utils.**
```

```
assets/
    logo.**
    background.**
```

```
config/
    settings.**
    db_config.**
```

---

## Notes

- The `src/` folder contains the main source code for the project.
- The `data/` folder is where Aruba Cloud Flavors (kaas, cloudservers, etc.) are stored.
- The `docs/` folder contains documentation files.
- The `tests/` folder includes all unit and integration tests.
- The `assets/` folder is used for storing images and other media files.
- The `config/` folder holds various configuration files.
