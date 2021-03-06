from typing import Any, Dict, Optional, Union, cast

import httpx

from ...client import AuthenticatedClient
from ...types import UNSET, Response


def _get_kwargs(
    *,
    client: AuthenticatedClient,
    name: str,
    content: str,
    initial_group: int,
) -> Dict[str, Any]:
    url = "{}/macro/create".format(client.base_url)

    headers: Dict[str, str] = client.get_headers()
    cookies: Dict[str, Any] = client.get_cookies()

    params: Dict[str, Any] = {}
    params["name"] = name

    params["content"] = content

    params["initialGroup"] = initial_group

    params = {k: v for k, v in params.items() if v is not UNSET and v is not None}

    return {
        "method": "post",
        "url": url,
        "headers": headers,
        "cookies": cookies,
        "timeout": client.get_timeout(),
        "params": params,
    }


def _parse_response(*, response: httpx.Response) -> Optional[Union[Any, int]]:
    if response.status_code == 200:
        response_200 = cast(int, response.json())
        return response_200
    if response.status_code == 400:
        response_400 = cast(Any, None)
        return response_400
    if response.status_code == 405:
        response_405 = cast(Any, None)
        return response_405
    return None


def _build_response(*, response: httpx.Response) -> Response[Union[Any, int]]:
    return Response(
        status_code=response.status_code,
        content=response.content,
        headers=response.headers,
        parsed=_parse_response(response=response),
    )


def sync_detailed(
    *,
    client: AuthenticatedClient,
    name: str,
    content: str,
    initial_group: int,
) -> Response[Union[Any, int]]:
    """Create macro

    Args:
        name (str):
        content (str):
        initial_group (int):

    Returns:
        Response[Union[Any, int]]
    """

    kwargs = _get_kwargs(
        client=client,
        name=name,
        content=content,
        initial_group=initial_group,
    )

    response = httpx.request(
        verify=client.verify_ssl,
        **kwargs,
    )

    return _build_response(response=response)


def sync(
    *,
    client: AuthenticatedClient,
    name: str,
    content: str,
    initial_group: int,
) -> Optional[Union[Any, int]]:
    """Create macro

    Args:
        name (str):
        content (str):
        initial_group (int):

    Returns:
        Response[Union[Any, int]]
    """

    return sync_detailed(
        client=client,
        name=name,
        content=content,
        initial_group=initial_group,
    ).parsed


async def asyncio_detailed(
    *,
    client: AuthenticatedClient,
    name: str,
    content: str,
    initial_group: int,
) -> Response[Union[Any, int]]:
    """Create macro

    Args:
        name (str):
        content (str):
        initial_group (int):

    Returns:
        Response[Union[Any, int]]
    """

    kwargs = _get_kwargs(
        client=client,
        name=name,
        content=content,
        initial_group=initial_group,
    )

    async with httpx.AsyncClient(verify=client.verify_ssl) as _client:
        response = await _client.request(**kwargs)

    return _build_response(response=response)


async def asyncio(
    *,
    client: AuthenticatedClient,
    name: str,
    content: str,
    initial_group: int,
) -> Optional[Union[Any, int]]:
    """Create macro

    Args:
        name (str):
        content (str):
        initial_group (int):

    Returns:
        Response[Union[Any, int]]
    """

    return (
        await asyncio_detailed(
            client=client,
            name=name,
            content=content,
            initial_group=initial_group,
        )
    ).parsed
